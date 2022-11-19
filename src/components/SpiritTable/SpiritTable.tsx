import React, {FC} from 'react';
import Table from 'react-bootstrap/Table';

import {Spirit, SpiritAction} from '../../lib/api/spirits/v1/spirit.pb';

interface SpiritTableProps {
  spirits: Spirit[]
}

const SpiritTable: FC<SpiritTableProps> = (props) => {
  return (
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {props.spirits.map((spirit: Spirit, i: number) =>
          <tr key={i}>
            <td>{spirit.meta?.id}</td>
            <td>{spirit.name}</td>
            <td>{spirit.actions?.map((action: SpiritAction) =>
              action.name + ' ')}</td>
          </tr>,
        )}
      </tbody>
    </Table>
  );
};

export default SpiritTable;
