const action = require('./action');
const battle = require('./battle');

// Remove me when https://github.com/facebook/jest/issues/12628 is fixed.
const structuredClone = (obj) => {
  const clone = JSON.parse(JSON.stringify(obj));
  clone.map((spirit) => spirit.action = action.attack);
  return clone;
};

describe('Battle', () => {
  let spirits, callbackSpirits, b;

  beforeEach(() => {
    spirits = [
      {name: 'a', stats: {health: 4, power: 1}, action: action.attack},
      {name: 'b', stats: {health: 4, power: 2}, action: action.attack},
    ];
    callbackSpirits = [];
    b = new battle.Battle(structuredClone(spirits), jest.fn((spirits) => {
      callbackSpirits.push(structuredClone(spirits));
    }));
  });

  describe('start', () => {
    beforeEach(() => {
      b.start();
    });

    it('runs the battle and calls the callback', () => {
      const wantSpirits = structuredClone(spirits);
      const wantSpiritsA = wantSpirits[0];
      const wantSpiritsB = wantSpirits[1];
      const wantCallbackSpirits = [];

      // Initial callback with initial spirits.
      wantCallbackSpirits.push(structuredClone(wantSpirits));
      
      // Spirit a goes first.
      wantSpiritsA.action(wantSpiritsA, wantSpiritsB);
      wantCallbackSpirits.push(structuredClone(wantSpirits));

      // Spirit b goes next.
      wantSpiritsB.action(wantSpiritsB, wantSpiritsA);
      wantCallbackSpirits.push(structuredClone(wantSpirits));

      // Both go one more time, and then spirit a is knocked out.
      wantSpiritsA.action(wantSpiritsA, wantSpiritsB);
      wantCallbackSpirits.push(structuredClone(wantSpirits));
      wantSpiritsB.action(wantSpiritsB, wantSpiritsA);
      wantCallbackSpirits.push(structuredClone(wantSpirits));

      expect(callbackSpirits).toEqual(wantCallbackSpirits);
    });
  });
});