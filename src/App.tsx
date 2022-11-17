import React, {FC, useState, useEffect} from 'react';
import {Container, Nav, Navbar, Spinner} from 'react-bootstrap';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import './App.css';

import ActionTable from './components/ActionTable/ActionTable';
import BattleTable from './components/BattleTable/BattleTable';
import SpiritTable from './components/SpiritTable/SpiritTable';

import {Action} from './lib/api/spirits/v1/action.pb';
import {Battle} from './lib/api/spirits/v1/battle.pb';
import {Spirit} from './lib/api/spirits/v1/spirit.pb';

interface BattleClient {
  createBattle(): Promise<Battle>
  listBattles(): Promise<Battle[]>
};

interface SpiritClient {
  listSpirits(): Promise<Spirit[]>
};

interface ActionClient {
  listActions(): Promise<Action[]>
};

interface AppProps {
  battleClient: BattleClient
  spiritClient: SpiritClient
  actionClient: ActionClient
};

const App: FC<AppProps> = (props) => {
  const [battles, setBattles] = useState<Battle[]>([]);
  const [spirits, setSpirits] = useState<Spirit[]>([]);
  const [actions, setActions] = useState<Action[]>([]);

  useEffect(() => {
    props.battleClient
        .listBattles()
        .then(setBattles)
        .catch((error) => {
          console.error(`list battles: ${error.toString()}`);
        });
  }, [props.battleClient]);

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

  const onNewBattle = () => {
    props.battleClient.createBattle()
        .then((battle) => {
          console.log(`battle: ${battle}`);
        }).catch((error) => {
          console.error(`create battle: ${error.toString()}`);
        });
  };

  return (
    <div className="font-monospace p-3">
      <Navbar bg="light" className="mb-3">
        <Container>
          <Navbar.Brand href="/">Spirits</Navbar.Brand>
          <Nav>
            <Nav.Link onClick={onNewBattle}>New Battle</Nav.Link>
          </Nav>
        </Container>
      </Navbar>

      <Tabs
        defaultActiveKey="battles"
        id="uncontrolled-tab-example"
        className="mb-3"
      >
        <Tab eventKey="battles" title="Battles">
          {battles.length > 0 ?
            <BattleTable battles={battles} /> :
            <Spinner animation="border" role="status" />}
        </Tab>
        <Tab eventKey="spirits" title="Spirits">
          {battles.length > 0 ?
            <SpiritTable spirits={spirits} /> :
            <Spinner animation="border" role="status" />}
        </Tab>
        <Tab eventKey="actions" title="Actions">
          {battles.length > 0 ?
            <ActionTable actions={actions} /> :
            <Spinner animation="border" role="status" />}
        </Tab>
      </Tabs>
    </div>
  );
};

export default App;
