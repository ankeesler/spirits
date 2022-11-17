import React from 'react';
import {render} from '@testing-library/react';
import App from './App';

import {FakeSpiritClient} from './lib/client/spirit';
import {FakeActionClient} from './lib/client/action';

test('renders learn react link', () => {
  render(<App
    spiritClient={new FakeSpiritClient([])}
    actionClient={new FakeActionClient([])} />);
});
