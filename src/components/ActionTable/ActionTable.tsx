import React, {FC} from 'react';
import Table from 'react-bootstrap/Table';

import {Action} from '../../lib/api/spirits/v1/action.pb';

interface ActionTableProps {
  actions: Action[]
}

const ActionTable: FC<ActionTableProps> = (props) => {
  return (
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>ID</th>
          <th>Description</th>
        </tr>
      </thead>
      <tbody>
        {props.actions.map((Action: Action, i: number) =>
          <tr key={i}>
            <td>{Action.meta?.id}</td>
            <td>{Action.description}</td>
          </tr>,
        )}
      </tbody>
    </Table>
  );
};

export default ActionTable;
