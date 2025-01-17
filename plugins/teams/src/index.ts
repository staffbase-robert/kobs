import { IPluginComponents } from '@kobsio/plugin-core';

import icon from './assets/icon.png';

import Home from './components/home/Home';
import Page from './components/page/Page';
import Panel from './components/panel/Panel';

import TeamsItem from './components/page/TeamsItem';

const teamsPlugin: IPluginComponents = {
  teams: {
    home: Home,
    icon: icon,
    page: Page,
    panel: Panel,
  },
};

export default teamsPlugin;

export { TeamsItem };
