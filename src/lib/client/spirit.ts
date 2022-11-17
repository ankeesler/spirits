import {Spirit, SpiritService} from '../api/spirits/v1/spirit.pb';

const INIT_REQ = {pathPrefix: '/api'};

export class SpiritClient {
  async listSpirits(): Promise<Spirit[]> {
    // eslint-disable-next-line new-cap
    return (await SpiritService.ListSpirits({}, INIT_REQ)).spirits!;
  }
};

export class FakeSpiritClient {
  private spirits: Spirit[];

  constructor(Spirits: Spirit[]) {
    this.spirits = Spirits;
  }

  listSpirits(): Promise<Spirit[]> {
    return Promise.resolve(this.spirits);
  }
};
