import React, {FC, FormEvent, useEffect, useState} from 'react';
import {Container, Spinner} from 'react-bootstrap';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
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
  const [battle, setBattle] = useState<Battle | undefined>(undefined);
  const [showAddTeamModal, setShowAddTeamModal] = useState<boolean>(false);
  const [teamName, setTeamName] = useState<string>('');

  useEffect(() => {
    props.battleClient.watchBattle(id as string, (battle: Battle) => {
      console.log(`BattleView: got battle: ${battle.toString()}`);
      setBattle(battle);
    });
  }, []);

  const onAddTeam = (e: FormEvent) => {
    e.preventDefault();

    props.battleClient.addTeam(battle!.meta!.id!, teamName)
        .then(setBattle)
        .catch((error) => {
          alert(error.message);
        });

    setShowAddTeamModal(false);
    setTeamName('');
  };

  return (
    <Container>
      <Button className="mb-3" onClick={() => setShowAddTeamModal(true)}>
        Add Team
      </Button>

      {battle ? <BattleCard battle={battle!} /> : <Spinner />}

      <Modal show={showAddTeamModal} onHide={() => setShowAddTeamModal(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Add team</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={onAddTeam}>
            <Form.Group className="mb-3">
              <Form.Label>Team name</Form.Label>
              <Form.Control
                value={teamName}
                onChange={(e) => setTeamName(e.target.value)}
                autoFocus />
            </Form.Group>
            <Button type="submit">Add</Button>
          </Form>
        </Modal.Body>
      </Modal>
    </Container>
  );
};

export default BattleView;
