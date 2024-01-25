import React from 'react';
import ReactDOM from 'react-dom';
import './styles.css';
import App from './App';

const root = document.getElementById('root');

// Use createRoot instead of ReactDOM.render
const reactRoot = ReactDOM.createRoot(root);

reactRoot.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
