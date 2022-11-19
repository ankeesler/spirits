import React, {FC, useEffect, useState} from 'react';
import {Container, Spinner} from 'react-bootstrap';
import {useLoaderData} from 'react-router-dom';
import BattleCard from '../components/BattleCard/BattleCard';

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
  const [loaded, setLoaded] = useState<Boolean>(false);

  useEffect(() => {
    props.battleClient.watchBattle(id as string, (battle: Battle) => {
      console.log(`BattleView: got battle: ${battle.toString()}`);
      setLoaded(true);
      setBattle(battle);
    });
  }, []);

  return (
    <Container>
      {loaded ? <BattleCard battle={battle} /> : <Spinner />}
    </Container>
  );
};

export default BattleView;
