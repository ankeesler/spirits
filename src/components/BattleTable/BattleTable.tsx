import React, {FC} from 'react';
import Table from 'react-bootstrap/Table';
import './BattleTable.css';

import {Battle} from '../../lib/api/spirits/v1/battle.pb';

interface BattleTableProps {
  battles: Battle[]
}

const BattleTable: FC<BattleTableProps> = (props) => {
  return (
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>ID</th>
          <th>Created</th>
          <th>Updated</th>
          <th>State</th>
        </tr>
      </thead>
      <tbody>
        {props.battles.map((battle: Battle, i: number) =>
          <tr key={i}>
            <td>{battle.meta?.id}</td>
            <td>{battle.meta?.createdTime}</td>
            <td>{battle.meta?.updatedTime}</td>
            <td>{battle.state}</td>
          </tr>,
        )}
      </tbody>
    </Table>
  );
};

export default BattleTable;
