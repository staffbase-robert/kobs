import {
  Alert,
  AlertVariant,
  Drawer,
  DrawerContent,
  DrawerContentBody,
  PageSection,
  PageSectionVariants,
  Title,
} from '@patternfly/react-core';
import React, { useMemo, useState } from 'react';
import { useHistory, useLocation } from 'react-router-dom';

import ApplicationsToolbar from './ApplicationsToolbar';
import { IOptions } from '../../utils/interfaces';
import Panel from '../panel/Panel';
import { getInitialOptions } from '../../utils/helpers';

export interface IApplicationsProps {
  name: string;
  displayName: string;
  description: string;
}

// Applications is the page which lets the user query all the created applications by cluster and namespace. The user
// can also select the view he wants to see (gallery vs. topology). The component is just a wrapper for the toolbar and
// the panel. It handles the reflection of the selected clusters and namespaces in the current url.
const Applications: React.FunctionComponent<IApplicationsProps> = ({
  name,
  displayName,
  description,
}: IApplicationsProps) => {
  const history = useHistory();
  const location = useLocation();
  const [options, setOptions] = useState<IOptions>(useMemo<IOptions>(() => getInitialOptions(), []));
  const [selectedApplication, setSelectedApplication] = useState<React.ReactNode>(undefined);

  // changeOptions is used to change the options. Besides setting a new value for the options state we also reflect the
  // options in the current url.
  const changeOptions = (opts: IOptions): void => {
    const c = opts.clusters.map((cluster) => `&cluster=${cluster}`);
    const n = opts.namespaces.map((namespace) => `&namespace=${namespace}`);

    history.push({
      pathname: location.pathname,
      search: `?view=${opts.view}${c.length > 0 ? c.join('') : ''}${n.length > 0 ? n.join('') : ''}`,
    });

    setOptions(opts);
  };

  return (
    <React.Fragment>
      <PageSection variant={PageSectionVariants.light}>
        <Title headingLevel="h6" size="xl">
          {displayName}
        </Title>
        <p>{description}</p>
        <ApplicationsToolbar options={options} setOptions={changeOptions} />
      </PageSection>

      <Drawer isExpanded={selectedApplication !== undefined}>
        <DrawerContent panelContent={selectedApplication}>
          <DrawerContentBody>
            <PageSection
              style={options.view === 'topology' ? { height: '100%', minHeight: '100%' } : { minHeight: '100%' }}
              variant={PageSectionVariants.default}
            >
              {options.clusters.length === 0 ? (
                <Alert variant={AlertVariant.info} title="You have to select at least one cluster">
                  <p>Select a list of clusters and namespaces from the toolbar.</p>
                </Alert>
              ) : (
                <Panel
                  defaults={{ cluster: '', name: '', namespace: '' }}
                  name={name}
                  title=""
                  options={{
                    clusters: options.clusters,
                    namespaces: options.namespaces,
                    team: undefined,
                    view: options.view,
                  }}
                  times={options.times}
                  showDetails={setSelectedApplication}
                />
              )}
            </PageSection>
          </DrawerContentBody>
        </DrawerContent>
      </Drawer>
    </React.Fragment>
  );
};

export default Applications;
