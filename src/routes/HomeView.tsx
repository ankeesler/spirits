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

interface HomeViewProps {
  battleClient: BattleClient
  spiritClient: SpiritClient
  actionClient: ActionClient
};

const HomeView: FC<HomeViewProps> = (props) => {
  const [battles, setBattles] = useState<Battle[]>([]);
  const [battlesLoaded, setBattlesLoaded] = useState<Boolean>(false);
  const [spirits, setSpirits] = useState<Spirit[]>([]);
  const [spiritsLoaded, setSpiritsLoaded] = useState<Boolean>(false);
  const [actions, setActions] = useState<Action[]>([]);
  const [actionsLoaded, setActionsLoaded] = useState<Boolean>(false);

  useEffect(() => {
    props.battleClient
        .listBattles()
        .then((battles: Battle[]) => {
          setBattlesLoaded(true);
          setBattles(battles);
        }).catch((error) => {
          console.error(`list battles: ${error.toString()}`);
        });
  }, [props.battleClient]);

  useEffect(() => {
    props.spiritClient
        .listSpirits()
        .then((spirits: Spirit[]) => {
          setSpiritsLoaded(true);
          setSpirits(spirits);
        }).catch((error) => {
          console.error(`list spirits: ${error.toString()}`);
        });
  }, [props.spiritClient]);

  useEffect(() => {
    props.actionClient.listActions()
        .then((actions: Action[]) => {
          setActionsLoaded(true);
          setActions(actions);
        }).catch((error) => {
          console.error(`list actions: ${error.toString()}`);
        });
  }, [props.actionClient]);

  return <Tabs
    defaultActiveKey="battles"
    id="uncontrolled-tab-example"
    className="mb-3"
  >
    <Tab eventKey="battles" title="Battles">
      {battlesLoaded ?
        <BattleTable battles={battles} /> :
        <Spinner animation="border" role="status" />}
    </Tab>
    <Tab eventKey="spirits" title="Spirits">
      {spiritsLoaded ?
        <SpiritTable spirits={spirits} /> :
        <Spinner animation="border" role="status" />}
    </Tab>
    <Tab eventKey="actions" title="Actions">
      {actionsLoaded ?
        <ActionTable actions={actions} /> :
        <Spinner animation="border" role="status" />}
    </Tab>
  </Tabs>;
};

export default HomeView;
