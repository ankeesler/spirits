import React from 'react';
import {render} from '@testing-library/react';
import App from './App';

import {FakeBattleClient} from './lib/client/battle';
import {FakeSpiritClient} from './lib/client/spirit';
import {FakeActionClient} from './lib/client/action';

test('renders learn react link', () => {
  render(<App
    battleClient={new FakeBattleClient([])}
    spiritClient={new FakeSpiritClient([])}
    actionClient={new FakeActionClient([])} />);
});
