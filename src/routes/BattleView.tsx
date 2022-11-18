import React, {FC} from 'react';
import {Spinner} from 'react-bootstrap';
import {useLoaderData} from 'react-router-dom';

import {Battle} from '../lib/api/spirits/v1/battle.pb';

interface BattleClient {
  addTeam(battleId: string, teamName: string): Promise<Battle>;
};

interface BattleViewProps {
  battleClient: BattleClient
};

const BattleView: FC<BattleViewProps> = (props) => {
  const battle = useLoaderData();
  console.log(`loaded battle: ${battle}`);
  return <Spinner animation="border" role="status" />;
};

export default BattleView;
