import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import ActionTable from './ActionTable';

describe('<ActionTable />', () => {
  test('it should mount', () => {
    render(<ActionTable actions={[]} />);
    const actionTable = screen.getByTestId('ActionTable');
    expect(actionTable).toBeInTheDocument();
  });
});