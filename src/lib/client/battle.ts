import {Battle, BattleService, BattleState} from '../api/spirits/v1/battle.pb';

const INIT_REQ = {pathPrefix: '/api'};

export class BattleClient {
  async createBattle(): Promise<Battle> {
    // eslint-disable-next-line new-cap
    return (await BattleService.CreateBattle({}, INIT_REQ)).battle!;
  }

  async listBattles(): Promise<Battle[]> {
    // eslint-disable-next-line new-cap
    return (await BattleService.ListBattles({}, INIT_REQ)).battles!;
  }
};

export class FakeBattleClient {
  private battles: Battle[];

  constructor(battles: Battle[]) {
    this.battles = battles;
  }

  createBattle(): Promise<Battle> {
    const now = new Date().toString();
    const battle = {
      meta: {
        id: 'ghi789',
        createdTime: now,
        updatedTimed: now,
      },
      state: BattleState.BATTLE_STATE_PENDING,
    };
    this.battles.push(battle);
    return Promise.resolve(battle);
  }

  listBattles(): Promise<Battle[]> {
    return Promise.resolve(this.battles);
  }
};
