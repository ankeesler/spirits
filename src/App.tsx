import React, {FC, useState, useEffect} from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import './App.css';

import ActionTable from './components/ActionTable/ActionTable';
import SpiritTable from './components/SpiritTable/SpiritTable';

import {Action} from './lib/api/spirits/v1/action.pb';
import {Spirit} from './lib/api/spirits/v1/spirit.pb';

interface SpiritClient {
  listSpirits(): Promise<Spirit[]>
};

interface ActionClient {
  listActions(): Promise<Action[]>
};

interface AppProps {
  spiritClient: SpiritClient
  actionClient: ActionClient
};

const App: FC<AppProps> = (props) => {
  const [spirits, setSpirits] = useState<Spirit[]>([]);
  const [actions, setActions] = useState<Action[]>([]);

  useEffect(() => {
    props.spiritClient
        .listSpirits()
        .then(setSpirits)
        .catch((error) => {
          console.error(`list spirits: ${error.toString()}`);
        });
  }, [props.spiritClient]);

  useEffect(() => {
    props.actionClient.listActions()
        .then(setActions)
        .catch((error) => {
          console.error(`list spirits: ${error.toString()}`);
        });
  }, [props.actionClient]);

  return (
    <Tabs
      defaultActiveKey="spirits"
      id="uncontrolled-tab-example"
      className="mb-3"
    >
      <Tab eventKey="spirits" title="Spirits">
        <SpiritTable spirits={spirits} />
      </Tab>
      <Tab eventKey="actions" title="Actions">
        <ActionTable actions={actions} />
      </Tab>
    </Tabs>
  );
};

export default App;
