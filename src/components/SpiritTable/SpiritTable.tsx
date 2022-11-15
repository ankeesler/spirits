import { useState, FC, useEffect } from 'react'
import Table from 'react-bootstrap/Table';
import './SpiritTable.css';
import { Spirit, SpiritAction, SpiritService } from '../../lib/api/spirits/v1/spirit.pb';

interface SpiritTablePros {
}

const SpiritTable: FC<SpiritTablePros> = () => {
  const [spirits, setSpirits] = useState<Spirit[]>([]);

  useEffect(() => {
    SpiritService.ListSpirits({}, {pathPrefix: '/api'}).then((rsp) => {
      if (rsp.spirits) {
        setSpirits(rsp.spirits!);
      }
    }).catch((error) => {
      console.error(`list spirits: ${error.toString()}`);
    });
  });

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
        {spirits.map((spirit: Spirit) => 
          <tr>
            <td>{spirit.meta?.id}</td>
            <td>{spirit.name}</td>
            <td>{spirit.actions?.map((action: SpiritAction) => action.name)}</td>
          </tr>
        )}
      </tbody>
    </Table>
  )
}

export default SpiritTable;
