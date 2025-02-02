import { Route, Switch } from 'react-router-dom';
import React from 'react';

import Application from './Application';
import Applications from './Applications';
import { IPluginOptions } from '../../utils/interfaces';
import { IPluginPageProps } from '@kobsio/plugin-core';

const Page: React.FunctionComponent<IPluginPageProps> = ({
  name,
  displayName,
  description,
  options,
}: IPluginPageProps) => {
  if (!options || !options.hasOwnProperty('prometheus')) {
    return null;
  }

  const pluginOptions: IPluginOptions = {
    klogs: options['klogs'],
    prometheus: options['prometheus'],
  };

  return (
    <Switch>
      <Route exact={true} path={`/${name}`}>
        <Applications name={name} displayName={displayName} description={description} pluginOptions={pluginOptions} />
      </Route>
      <Route exact={true} path={`/${name}/:namespace/:application`}>
        <Application name={name} pluginOptions={pluginOptions} />
      </Route>
    </Switch>
  );
};

export default Page;
