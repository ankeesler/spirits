import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import 'bootstrap/dist/css/bootstrap.min.css';
import { FakeSpiritClient, SpiritClient } from './lib/client/spirit';
import { ActionClient, FakeActionClient } from './lib/client/action';
import { Spirit } from './lib/api/spirits/v1/spirit.pb';
import { Action } from './lib/api/spirits/v1/action.pb';

const fakeSpirits: Spirit[] = [
  {
    meta: {
      id: 'abc123',
    },
    name: 'tuna',
    actions: [
      {
        name: 'flop',
      },
      {
        name: 'swim',
      },
    ],
  },
  {
    meta: {
      id: 'def456',
    },
    name: 'fish',
    actions: [
      {
        name: 'breathe',
      },
      {
        name: 'egg',
      },
    ],
  },
];

const fakeActions: Action[] = [
  {
    meta: {
      id: 'abc123',
    },
    description: 'This is a big one',
  },
  {
    meta: {
      id: 'def456',
    },
    description: 'Oh lawd he coming',
  },
];

const isDev = !process.env.NODE_ENV || process.env.NODE_ENV === 'development';
const spiritClient = isDev ? new FakeSpiritClient(fakeSpirits) : new SpiritClient();
const actionClient = isDev ? new FakeActionClient(fakeActions) : new ActionClient();

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <App spiritClient={spiritClient} actionClient={actionClient}/>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
