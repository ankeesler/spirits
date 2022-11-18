import {Battle, BattleService, BattleState} from '../api/spirits/v1/battle.pb';

const INIT_REQ = {pathPrefix: '/api'};

export type BattleCallback = (battle: Battle) => void

export class BattleClient {
  async createBattle(): Promise<Battle> {
    // eslint-disable-next-line new-cap
    return (await BattleService.CreateBattle({}, INIT_REQ)).battle!;
  }

  async listBattles(): Promise<Battle[]> {
    // eslint-disable-next-line new-cap
    return (await BattleService.ListBattles({}, INIT_REQ)).battles!;
  }

  async addTeam(battleId: string, teamName: string): Promise<Battle> {
    // eslint-disable-next-line new-cap
    return (await BattleService.AddBattleTeam({
      battleId: battleId,
      teamName: teamName,
    })).battle!;
  }

  async watchBattle(id: string, callback: BattleCallback): Promise<void> {
    // eslint-disable-next-line new-cap
    return BattleService.WatchBattle({id: id}, (rsp) => {
      callback(rsp.battle!);
    });
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

  addTeam(_: string, teamName: string): Promise<Battle> {
    this.battles[0].teams?.push({name: teamName});
    return Promise.resolve(this.battles[0]);
  }

  watchBattle(id: string, callback: BattleCallback): Promise<void> {
    setInterval(() => {
      callback(this.battles[0]);
    }, 1000);
    return Promise.resolve();
  }
};
