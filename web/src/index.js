import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';
import FakeClient from './lib/FakeClient'
import RealClient from './lib/RealClient'

import './index.css';

const scheme = (window.location.protocol === 'https:' ? 'wss://' : 'ws://');
const ws = new WebSocket(scheme + window.location.host + '/api/battle');

ReactDOM.render(
  <React.StrictMode>
    <App client={process.env.NODE_ENV === 'development' ? new FakeClient() : new RealClient(ws)}/>
  </React.StrictMode>,
  document.getElementById('root')
);
