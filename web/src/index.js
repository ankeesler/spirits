import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';
import FakeClient from './lib/FakeClient'
import RealClient from './lib/RealClient'

import './index.css';

ReactDOM.render(
  <React.StrictMode>
    <App client={process.env.NODE_ENV === 'development' ? new FakeClient() : new RealClient()}/>
  </React.StrictMode>,
  document.getElementById('root')
);
