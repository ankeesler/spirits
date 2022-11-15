import React from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import './App.css';

import ActionTable from './components/ActionTable/ActionTable';
import SpiritTable from './components/SpiritTable/SpiritTable';

function App() {
  return (
    <Tabs
      defaultActiveKey="spirits"
      id="uncontrolled-tab-example"
      className="mb-3"
    >
      <Tab eventKey="spirits" title="Spirits">
        <SpiritTable />
      </Tab>
      <Tab eventKey="actions" title="Actions">
        <ActionTable />
      </Tab>
    </Tabs>
  );
}

export default App;
