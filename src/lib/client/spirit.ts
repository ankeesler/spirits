import { Spirit, SpiritService } from "../api/spirits/v1/spirit.pb";

const INIT_REQ = {pathPrefix: '/api'}

export class SpiritClient {
  listSpirits(): Promise<Spirit[]> {
    return SpiritService.ListSpirits({}, INIT_REQ).then((rsp) => {
      return rsp.spirits!;
    }).catch((error) => {
      return error;
    });
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
