import { IRow, Table, TableBody, TableHeader } from '@patternfly/react-table';
import React, { useContext } from 'react';
import { useQuery } from 'react-query';

import { ClustersContext, IClusterContext, emptyState } from '@kobsio/plugin-core';
import Details from '../panel/details/Details';
import { TApiType } from '../../utils/interfaces';

interface IPanelListProps {
  name: string;
  title: string;
  cluster: string;
  namespace: string;
  selector?: string;
  type: TApiType;
  showDetails?: (details: React.ReactNode) => void;
}

const PanelList: React.FunctionComponent<IPanelListProps> = ({
  name,
  title,
  type,
  cluster,
  namespace,
  selector,
  showDetails,
}: IPanelListProps) => {
  const clustersContext = useContext<IClusterContext>(ClustersContext);
  const resource =
    clustersContext.resources && clustersContext.resources.hasOwnProperty(type)
      ? clustersContext.resources[type]
      : undefined;

  const { isError, isLoading, error, data, refetch } = useQuery<IRow[], Error>(
    ['flux/list', name, cluster, type, namespace, selector],
    async () => {
      try {
        if (!resource) {
          throw new Error('Could not find resource');
        }

        const response = await fetch(
          `/api/plugins/resources/resources?cluster=${cluster}&namespace=${namespace}${
            selector ? `&paramName=labelSelector&param=${selector}` : ''
          }&resource=${resource.resource}&path=/apis/${resource.path}`,
          { method: 'get' },
        );
        const json = await response.json();

        if (response.status >= 200 && response.status < 300) {
          return resource.rows(json);
        } else {
          if (json.error) {
            throw new Error(json.error);
          } else {
            throw new Error('An unknown error occured');
          }
        }
      } catch (err) {
        throw err;
      }
    },
  );

  // refetchhWithDelay is used to call the refetch function to get the resource, but with a delay of 3 seconds. This is
  // required, because sometime the Kubenretes isn't that fast after an action (edit, delete, ...) was triggered.
  const refetchhWithDelay = (): void => {
    setTimeout(() => {
      refetch();
    }, 3000);
  };

  return (
    <Table
      aria-label={title}
      variant="compact"
      borders={false}
      cells={resource?.columns || ['']}
      rows={
        data && data.length > 0 && data[0].cells?.length === resource?.columns.length
          ? data
          : emptyState(resource?.columns.length || 3, isLoading, isError, error)
      }
    >
      <TableHeader />
      <TableBody
        onRowClick={
          showDetails && resource && data && data.length > 0 && data[0].cells?.length === resource.columns.length
            ? (e, row, props, data): void =>
                showDetails(
                  <Details
                    name={name}
                    type={type}
                    request={resource}
                    resource={row}
                    close={(): void => showDetails(undefined)}
                    refetch={refetchhWithDelay}
                  />,
                )
            : undefined
        }
      />
    </Table>
  );
};

export default PanelList;
