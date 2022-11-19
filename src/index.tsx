import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  createBrowserRouter,
  RouterProvider,
} from 'react-router-dom';
import reportWebVitals from './reportWebVitals';
import 'bootstrap/dist/css/bootstrap.min.css';

import RootView from './routes/RootView';
import HomeView from './routes/HomeView';
import BattleView from './routes/BattleView';
import ErrorView from './routes/ErrorView';

import {FakeBattleClient, BattleClient} from './lib/client/battle';
import {FakeSpiritClient, SpiritClient} from './lib/client/spirit';
import {FakeActionClient, ActionClient} from './lib/client/action';
import {Battle, BattleState} from './lib/api/spirits/v1/battle.pb';
import {Spirit} from './lib/api/spirits/v1/spirit.pb';
import {Action} from './lib/api/spirits/v1/action.pb';

const fakeBattles: Battle[] = [
  {
    meta: {
      id: 'abc123',
      createdTime: '2022-10-10',
      updatedTime: '2022-10-10',
    },
    state: BattleState.BATTLE_STATE_FINISHED,
  },
  {
    meta: {
      id: 'def456',
      createdTime: '2022-11-16',
      updatedTime: '2022-11-18',
    },
    state: BattleState.BATTLE_STATE_WAITING,
  },
];

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
const battleClient =
  isDev ? new FakeBattleClient(fakeBattles) : new BattleClient();
const spiritClient =
  isDev ? new FakeSpiritClient(fakeSpirits) : new SpiritClient();
const actionClient =
  isDev ? new FakeActionClient(fakeActions) : new ActionClient();

const router = createBrowserRouter([
  {
    path: '/',
    element: <RootView
      battleClient={battleClient}
      spiritClient={spiritClient}
      actionClient={actionClient} />,
    errorElement: <ErrorView />,
    children: [
      {
        path: '/',
        element: <HomeView
          battleClient={battleClient}
          spiritClient={spiritClient}
          actionClient={actionClient} />,
      },
      {
        path: '/battles/:id',
        element: <BattleView battleClient={battleClient}/>,
        loader: async ({params}) => {
          return params.id;
        },
      },
    ],
  },
]);

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);
root.render(
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
