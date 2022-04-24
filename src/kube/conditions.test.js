const conditions = require('./conditions');

jest.mock('./date', () => {
  return () => 'some-date';
});

describe('conditions', () => {
  describe('upsert', () => {

    describe('no status', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('adds a new condition', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(true);
      })
    });

    describe('no conditions', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
          status: {},
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('adds a new condition', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(true);
      })
    });

    describe('non-matching conditions', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-other-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
            ],
          },
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('adds a new condition', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-other-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition except for status', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-other-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: 555,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('updates the existing condition', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition except for reason', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-other-reason',
                message: 'some-message',
                observedGeneration: 555,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('updates the existing condition', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition except for message', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-other-message',
                observedGeneration: 555,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('updates the existing condition', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition except for observed generation', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: 554,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('updates the existing condition', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: 555,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason', 'some-message');
      });

      it('does not update the object', () => {
        expect(obj).toEqual({
          metadata: {
            generation: 555,
          },
          status: {
            conditions: [
              {
                type: 'some-other-type',
                status: 'some-status',
                reason: 'some-reason',
                lastTransitionTime: 'some-date',
              },
              {
                type: 'some-type',
                status: 'some-status',
                reason: 'some-reason',
                message: 'some-message',
                observedGeneration: obj.metadata.generation,
                lastTransitionTime: 'some-date',
              },
            ],
          },
        })
        expect(updated).toEqual(false);
      })
    });
  });
});
