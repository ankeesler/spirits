import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import SpiritTable from './SpiritTable';

describe('<SpiritTable />', () => {
  test('it should mount', () => {
    render(<SpiritTable />);
    
    // const spiritTable = screen.getByTestId('SpiritTable');

    // expect(spiritTable).toBeInTheDocument();
  });
});