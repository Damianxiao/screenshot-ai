import React, { useState, useEffect, useRef } from 'react';
import ReactMarkdown from 'react-markdown';
import './App.css';

const App = () => {
  const [screenshot, setScreenshot] = useState('');
  const [model, setModel] = useState('gpt-4o');
  const [messages, setMessages] = useState([]);
  const ws = useRef(null);
  const outputRef = useRef(null);

  useEffect(() => {
    const wsUrl = process.env.REACT_APP_WS_URL || 'ws://localhost:3001/v1/ws';
    ws.current = new WebSocket(wsUrl);

    ws.current.onopen = () => {
      setMessages(prev => [...prev, { timestamp: new Date().toLocaleTimeString(), content: 'connected to server...' }]);
    };

    ws.current.onmessage = (event) => {
      // console.log('收到原始消息:', event.data);
      const msg = JSON.parse(event.data);
      // console.log('解析后消息:', msg);
      if (msg.type === 1) {
        setScreenshot(msg.content);

      } else if (msg.type === 3) {
        setMessages(prev => [...prev, { timestamp: new Date().toLocaleTimeString(), content: msg.content, isMarkdown: true }]);
      }
    };

    ws.current.onclose = () => {
      setScreenshot('');
      setMessages(prev => [...prev, { timestamp: new Date().toLocaleTimeString(), content: 'disconnect from server' }]);
    };

    return () => {
      if (ws.current && ws.current.readyState !== WebSocket.CLOSED) {
        ws.current.close();
      }
    };
  }, []);

  useEffect(() => {
    if (outputRef.current) {
      outputRef.current.scrollTop = outputRef.current.scrollHeight;
    }
  }, [messages]);

  const handleContinue = () => {
    if (!screenshot) {
      alert('wait screenshot');
      return;
    }
    if (ws.current.readyState !== WebSocket.OPEN) {
      alert('connect to server first');
      return;
    }
    const msg = {
      type: 2,
      content: screenshot,
      model: model,
    };
    ws.current.send(JSON.stringify(msg));
    setMessages(prev => [...prev, { timestamp: new Date().toLocaleTimeString(), content: `req sended,use model: ${model}` }]);
  };

  return (
    <div className="App">
      <div className="screenshot-container">
        {screenshot ? (
          <img src={`data:image/png;base64,${screenshot}`} alt="" className="screenshot"
          onClick={handleContinue} />
        ) : (
          <p className="placeholder">wait for screenshot...</p>
        )}
      </div>
      <div className="controls">
        <select value={model} onChange={(e) => setModel(e.target.value)} className="model-select">
          <option value="gpt-4o">GPT-4o</option>
          <option value="claude-3-5-sonnet-20240620">claude-3.5</option>
          <option value="gemini-1.5-flash">gemini-1.5</option>
          <option value="grok-2-vision-1212">grok-2</option>
        </select>
        <button onClick={handleContinue} className="continue-btn">answer</button>
      </div>
      <div className="output" ref={outputRef}>
        {messages.map((msg, index) => (
          <div key={index} className="message">
            <span className="timestamp">{msg.timestamp}</span>
            {msg.isMarkdown ? (
              <ReactMarkdown className="markdown">{msg.content}</ReactMarkdown>
            ) : (
              <span className="content">{msg.content}</span>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};

export default App;