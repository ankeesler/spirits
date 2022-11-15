import React, { FC } from 'react';
import './ActionTable.css';

interface ActionTableProps {}

const ActionTable: FC<ActionTableProps> = () => (
  <div className="ActionTable" data-testid="ActionTable">
    ActionTable Component
  </div>
);

export default ActionTable;
