const conditions = require('./conditions');

jest.mock('./date', () => {
  return () => 'some-date';
});

describe('conditions', () => {
  describe('upsert', () => {

    describe('no conditions', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {};
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason');
      });

      it('adds a new condition', () => {
        expect(obj).toEqual({
          conditions: [
            {
              type: 'some-type',
              status: 'some-status',
              reason: 'some-reason',
              lastTransitionTime: 'some-date',
            },
          ],
        })
        expect(updated).toEqual(true);
      })
    });

    describe('non-matching conditions', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
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
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason');
      });

      it('adds a new condition', () => {
        expect(obj).toEqual({
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
              lastTransitionTime: 'some-date',
            },
          ],
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition except for status', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
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
              lastTransitionTime: 'some-date',
            },
          ],
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason');
      });

      it('updates the existing condition', () => {
        expect(obj).toEqual({
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
              lastTransitionTime: 'some-date',
            },
          ],
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition except for reason', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
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
              lastTransitionTime: 'some-date',
            },
          ],
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason');
      });

      it('updates the existing condition', () => {
        expect(obj).toEqual({
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
              lastTransitionTime: 'some-date',
            },
          ],
        })
        expect(updated).toEqual(true);
      })
    });

    describe('matching condition', () => {
      let obj, updated;
      beforeEach(() => {
        obj = {
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
              lastTransitionTime: 'some-date',
            },
          ],
        };
        updated = conditions.upsert(obj, 'some-type', 'some-status', 'some-reason');
      });

      it('does not update the object', () => {
        expect(obj).toEqual({
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
              lastTransitionTime: 'some-date',
            },
          ],
        })
        expect(updated).toEqual(false);
      })
    });
  });
});
