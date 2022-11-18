import React, {FC, useEffect, useState} from 'react';
import {Spinner} from 'react-bootstrap';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';

import ActionTable from '../components/ActionTable/ActionTable';
import BattleTable from '../components/BattleTable/BattleTable';
import SpiritTable from '../components/SpiritTable/SpiritTable';

import {Action} from '../lib/api/spirits/v1/action.pb';
import {Battle} from '../lib/api/spirits/v1/battle.pb';
import {Spirit} from '../lib/api/spirits/v1/spirit.pb';

interface BattleClient {
  listBattles(): Promise<Battle[]>
};

interface SpiritClient {
  listSpirits(): Promise<Spirit[]>
};

interface ActionClient {
  listActions(): Promise<Action[]>
};

interface HomeProps {
  battleClient: BattleClient
  spiritClient: SpiritClient
  actionClient: ActionClient
};

const Home: FC<HomeProps> = (props) => {
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

  return <Tabs
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
      {spirits.length > 0 ?
        <SpiritTable spirits={spirits} /> :
        <Spinner animation="border" role="status" />}
    </Tab>
    <Tab eventKey="actions" title="Actions">
      {actions.length > 0 ?
        <ActionTable actions={actions} /> :
        <Spinner animation="border" role="status" />}
    </Tab>
  </Tabs>;
};

export default Home;
