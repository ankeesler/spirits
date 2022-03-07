import {render, screen} from '@testing-library/react';

import FakeClient from './lib/FakeClient';

import App from './App';

test('smoke test (renders the header)', () => {
  const client = new FakeClient();
  render(<App client={client} />);
  const header = screen.getByText(/spirits/i);
  expect(header).toBeInTheDocument();
});

