import React, {FC, useState} from 'react';
import {Container, Nav, Navbar} from 'react-bootstrap';
import {Link, Outlet, useNavigate} from 'react-router-dom';

import {Action} from '../lib/api/spirits/v1/action.pb';
import {Battle} from '../lib/api/spirits/v1/battle.pb';
import {Spirit} from '../lib/api/spirits/v1/spirit.pb';

interface BattleClient {
  createBattle(): Promise<Battle>
  listBattles(): Promise<Battle[]>
};

interface SpiritClient {
  listSpirits(): Promise<Spirit[]>
};

interface ActionClient {
  listActions(): Promise<Action[]>
};

interface AppProps {
  battleClient: BattleClient
  spiritClient: SpiritClient
  actionClient: ActionClient
};

// eslint-disable-next-line no-unused-vars
enum LogLevel {
  // eslint-disable-next-line no-unused-vars
  INFO = 'INFO',
  // eslint-disable-next-line no-unused-vars
  ERROR = 'ERROR',
};

interface LogLine {
  time: Date
  level: LogLevel
  message: string
};

const RootView: FC<AppProps> = (props) => {
  const [logLines, setLogLines] = useState<LogLine[]>([{
    time: new Date(),
    level: LogLevel.INFO,
    message: 'Start a battle by clicking "New Battle"',
  }]);
  const navigate = useNavigate();

  const logger = {
    info: (message: string) => {
      logLines.push({
        time: new Date(),
        level: LogLevel.INFO,
        message: message,
      });
      setLogLines(logLines);
    },
    error: (message: string) => {
      logLines.push({
        time: new Date(),
        level: LogLevel.ERROR,
        message: message,
      });
      setLogLines(logLines);
    },
  };

  const onNewBattle = () => {
    props.battleClient.createBattle()
        .then((battle) => {
          logger.info(`New battle: ${battle.meta?.id}`);
          navigate(`/battles/${battle.meta?.id}`);
        }).catch((error) => {
          logger.error(`create battle: ${error.toString()}`);
        });
  };

  const onLogin = () => {
    console.log('login not supported');
  };

  const getRecentLogLines = (): LogLine[] => {
    return logLines.length === 1 ? [logLines[0]] : logLines.slice(-2);
  };

  const getLogLineClassName = (logLine: LogLine): string => {
    return logLine.level == LogLevel.INFO ? 'text-secondary' : 'test-danger';
  };

  const getLogLineString = (logLine: LogLine): string => {
    return `[${logLine.time.toLocaleTimeString()}] ${logLine.message}`;
  };

  document.onkeyup = (e: KeyboardEvent) => {
    if (e.which === 'N'.charCodeAt(0)) {
      onNewBattle();
    }
  };

  return (
    <div className="font-monospace p-3">
      <Navbar bg="light" className="mb-3">
        <Container>
          <Navbar.Collapse className="justify-content-begin">
            <Nav.Link onClick={onNewBattle}>New Battle</Nav.Link>
          </Navbar.Collapse>
          <Navbar.Collapse className="justify-content-center">
            <Navbar.Brand><Link to={'/'}>Spirits</Link></Navbar.Brand>
          </Navbar.Collapse>
          <Navbar.Collapse className="justify-content-end">
            <Nav.Link onClick={onLogin}>Login</Nav.Link>
          </Navbar.Collapse>
        </Container>
      </Navbar>

      <Container>
        <div className="bg-light mb-3">
          {getRecentLogLines().map((logLine, i) =>
            <div key={i} className={getLogLineClassName(logLine)}>
              {getLogLineString(logLine)}
            </div>)
          }
        </div>
      </Container>

      <Outlet />
    </div>
  );
};

export default RootView;
