import { Action, ActionService } from "../api/spirits/v1/action.pb";

const INIT_REQ = {pathPrefix: '/api'}

export class ActionClient {
  listActions(): Promise<Action[]> {
    return ActionService.ListActions({}, INIT_REQ).then((rsp) => {
      return rsp.actions!;
    }).catch((error) => {
      return error;
    });
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
