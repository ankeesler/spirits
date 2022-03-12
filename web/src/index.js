import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';
import createClient from './lib/createClient';

import './index.css';

createClient((client) => {
  ReactDOM.render(
    <React.StrictMode>
      <App client={client} />
    </React.StrictMode>,
    document.getElementById('root')
  );
});

