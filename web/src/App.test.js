import {render, screen} from '@testing-library/react';

import App from './App';

test('smoke test (renders the header)', () => {
  const client = {
    startBattle() { },
  }
  render(<App client={client} />);
  const header = screen.getByText(/spirits/i);
  expect(header).toBeInTheDocument();
});

