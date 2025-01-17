import { IPluginComponents } from '@kobsio/plugin-core';

import icon from './assets/icon.png';

import Page from './components/page/Page';
import Panel from './components/panel/Panel';
import Preview from './components/preview/Preview';

const prometheusPlugin: IPluginComponents = {
  prometheus: {
    icon: icon,
    page: Page,
    panel: Panel,
    preview: Preview,
  },
};

export default prometheusPlugin;

export * from './utils/interfaces';
export * from './utils/helpers';
export * from './components/panel/Chart';
