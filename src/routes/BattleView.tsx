import React, {FC, useEffect, useState} from 'react';
import {Container, Spinner} from 'react-bootstrap';
import {useLoaderData} from 'react-router-dom';

import {Battle} from '../lib/api/spirits/v1/battle.pb';

type BattleCallback = (battle: Battle) => void

interface BattleClient {
  addTeam(battleId: string, teamName: string): Promise<Battle>;
  watchBattle(battleId: string, callback: BattleCallback): Promise<void>
};

interface BattleViewProps {
  battleClient: BattleClient
};

const BattleView: FC<BattleViewProps> = (props) => {
  const id = useLoaderData();
  const [battle, setBattle] = useState<Battle>({});

  useEffect(() => {
    props.battleClient.watchBattle(id as string, setBattle);
  }, []);

  const getBody = () => {
    if (!battle.meta) {
      return <Spinner animation="border" role="status" />;
    }

    return (<div>battle.meta.id</div>);
  };

  return <Container>{getBody()}</Container>;
};

export default BattleView;
