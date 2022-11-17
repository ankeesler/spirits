import {Action, ActionService} from '../api/spirits/v1/action.pb';

const INIT_REQ = {pathPrefix: '/api'};

export class ActionClient {
  async listActions(): Promise<Action[]> {
    // eslint-disable-next-line new-cap
    return (await ActionService.ListActions({}, INIT_REQ)).actions!;
  }
};

export class FakeActionClient {
  private actions: Action[];

  constructor(actions: Action[]) {
    this.actions = actions;
  }

  listActions(): Promise<Action[]> {
    return Promise.resolve(this.actions);
  }
};
