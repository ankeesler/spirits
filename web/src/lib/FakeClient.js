class FakeClient {
  startBattle(spirits, callback) {
    setTimeout(() => {
      let errorMessage = '';
      try {
        JSON.parse(spirits);
      } catch (error) {
        errorMessage = error.message;
      }
      callback(errorMessage, `> summary
  a: 3
  b: 3
> summary
  a: 3
  b: 2
> summary
  a: 1
  b: 2
> summary
  a: 1
  b: 1
> summary
  a: 0
  b: 1
`)
    }, 1000);
  };

  generateSpirits(callback) {
    setTimeout(() => {
      callback('', `[
  {"name": "a", "health": 3, "power": 1, "agility": 1, "actions": [""]},
  {"name": "b", "health": 3, "power": 2, "agility": 1}
]`);
    }, 1000);
  };
};

export default FakeClient;