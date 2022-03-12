import {render, screen} from '@testing-library/react';

import App from './App';

test('smoke test (renders the header)', () => {
  render(<App client={{}} />);
  const header = screen.getByText(/spirits/i);
  expect(header).toBeInTheDocument();
});

