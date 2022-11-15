import React, { useState, useEffect } from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import './App.css';

import ActionTable from './components/ActionTable/ActionTable';
import SpiritTable from './components/SpiritTable/SpiritTable';

import { Action, ActionService } from './lib/api/spirits/v1/action.pb';
import { Spirit, SpiritService } from './lib/api/spirits/v1/spirit.pb';

function App() {
  const [spirits, setSpirits] = useState<Spirit[]>([]);
  const [actions, setActions] = useState<Action[]>([]);

  useEffect(() => {
    SpiritService.ListSpirits({}, {pathPrefix: '/api'}).then((rsp) => {
      if (rsp.spirits) {
        setSpirits(rsp.spirits!);
      }
    }).catch((error) => {
      console.error(`list spirits: ${error.toString()}`);
    });
  });

  useEffect(() => {
    ActionService.ListActions({}, {pathPrefix: '/api'}).then((rsp) => {
      if (rsp.actions) {
        setActions(rsp.actions!);
      }
    }).catch((error) => {
      console.error(`list actions: ${error.toString()}`);
    });
  });

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
}

export default App;
