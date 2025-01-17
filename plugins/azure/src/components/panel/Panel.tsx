import React, { memo } from 'react';

import { IPluginPanelProps, PluginCard, PluginOptionsMissing } from '@kobsio/plugin-core';
import { IPanelOptions } from '../../utils/interfaces';
import { services } from '../../utils/services';

import CIContainerGroups from '../containerinstances/ContainerGroups';
import CIDetailsContainerGroup from '../containerinstances/DetailsContainerGroup';
import CIDetailsContainerGroupActions from '../containerinstances/DetailsContainerGroupActions';
import CIDetailsLogs from '../containerinstances/DetailsLogs';

import KSDetailsKubernetesService from '../kubernetesservices/DetailsKubernetesService';
import KSDetailsNodePoolsWrapper from '../kubernetesservices/DetailsNodePoolsWrapper';
import KSKubernetesServices from '../kubernetesservices/KubernetesServices';

import VMSSDetailsVirtualMachineScaleSets from '../virtualmachinescalesets/DetailsVirtualMachineScaleSets';
import VMSSDetailsVirtualMachines from '../virtualmachinescalesets/DetailsVirtualMachines';
import VMSSVirtualMachineScaleSets from '../virtualmachinescalesets/VirtualMachineScaleSets';

import Metric from '../metrics/Metric';

const providerCI = services['containerinstances'].provider;
const providerKS = services['kubernetesservices'].provider;
const providerVMSS = services['virtualmachinescalesets'].provider;

interface IPanelProps extends IPluginPanelProps {
  options?: IPanelOptions;
}

export const Panel: React.FunctionComponent<IPanelProps> = ({
  name,
  title,
  description,
  times,
  options,
  showDetails,
}: IPanelProps) => {
  // Panels for container services.
  if (
    options?.type &&
    options?.type === 'containerinstances' &&
    options.containerinstances &&
    options.containerinstances.type === 'list' &&
    options.containerinstances.resourceGroups
  ) {
    return (
      <PluginCard title={title} description={description} transparent={true}>
        <CIContainerGroups
          name={name}
          resourceGroups={options.containerinstances.resourceGroups}
          setDetails={showDetails}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'containerinstances' &&
    options.containerinstances &&
    options.containerinstances.type === 'details' &&
    options.containerinstances.resourceGroup &&
    options.containerinstances.containerGroup
  ) {
    return (
      <PluginCard
        title={title}
        description={description}
        actions={
          <CIDetailsContainerGroupActions
            name={name}
            resourceGroup={options.containerinstances.resourceGroup}
            containerGroup={options.containerinstances.containerGroup}
            isPanelAction={true}
          />
        }
      >
        <CIDetailsContainerGroup
          name={name}
          resourceGroup={options.containerinstances.resourceGroup}
          containerGroup={options.containerinstances.containerGroup}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'containerinstances' &&
    options.containerinstances &&
    options.containerinstances.type === 'logs' &&
    options.containerinstances.resourceGroup &&
    options.containerinstances.containerGroup &&
    options.containerinstances.containers
  ) {
    return (
      <PluginCard title={title} description={description}>
        <CIDetailsLogs
          name={name}
          resourceGroup={options.containerinstances.resourceGroup}
          containerGroup={options.containerinstances.containerGroup}
          containers={options.containerinstances.containers}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'containerinstances' &&
    options.containerinstances &&
    options.containerinstances.type === 'metrics' &&
    options.containerinstances.resourceGroup &&
    options.containerinstances.containerGroup &&
    options.containerinstances.metricNames &&
    options.containerinstances.aggregationType &&
    times
  ) {
    return (
      <PluginCard title={title} description={description}>
        <Metric
          name={name}
          resourceGroup={options.containerinstances.resourceGroup}
          provider={providerCI + options.containerinstances.containerGroup}
          metricNames={options.containerinstances.metricNames}
          aggregationType={options.containerinstances.aggregationType}
          times={times}
        />
      </PluginCard>
    );
  }

  // Panels for kubernetes services.
  if (
    options?.type &&
    options?.type === 'kubernetesservices' &&
    options.kubernetesservices &&
    options.kubernetesservices.type === 'list' &&
    options.kubernetesservices.resourceGroups
  ) {
    return (
      <PluginCard title={title} description={description} transparent={true}>
        <KSKubernetesServices
          name={name}
          resourceGroups={options.kubernetesservices.resourceGroups}
          setDetails={showDetails}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'kubernetesservices' &&
    options.kubernetesservices &&
    options.kubernetesservices.type === 'details' &&
    options.kubernetesservices.resourceGroup &&
    options.kubernetesservices.managedCluster
  ) {
    return (
      <PluginCard title={title} description={description}>
        <KSDetailsKubernetesService
          name={name}
          resourceGroup={options.kubernetesservices.resourceGroup}
          managedCluster={options.kubernetesservices.managedCluster}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'kubernetesservices' &&
    options.kubernetesservices &&
    options.kubernetesservices.type === 'nodePools' &&
    options.kubernetesservices.resourceGroup &&
    options.kubernetesservices.managedCluster
  ) {
    return (
      <PluginCard title={title} description={description}>
        <KSDetailsNodePoolsWrapper
          name={name}
          resourceGroup={options.kubernetesservices.resourceGroup}
          managedCluster={options.kubernetesservices.managedCluster}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'kubernetesservices' &&
    options.kubernetesservices &&
    options.kubernetesservices.type === 'metrics' &&
    options.kubernetesservices.resourceGroup &&
    options.kubernetesservices.managedCluster &&
    options.kubernetesservices.metricNames &&
    options.kubernetesservices.aggregationType &&
    times
  ) {
    return (
      <PluginCard title={title} description={description}>
        <Metric
          name={name}
          resourceGroup={options.kubernetesservices.resourceGroup}
          provider={providerKS + options.kubernetesservices.managedCluster}
          metricNames={options.kubernetesservices.metricNames}
          aggregationType={options.kubernetesservices.aggregationType}
          times={times}
        />
      </PluginCard>
    );
  }

  // Panel for virtual machine scale sets
  if (
    options?.type &&
    options?.type === 'virtualmachinescalesets' &&
    options.virtualmachinescalesets &&
    options.virtualmachinescalesets.type === 'list' &&
    options.virtualmachinescalesets.resourceGroups
  ) {
    return (
      <PluginCard title={title} description={description} transparent={true}>
        <VMSSVirtualMachineScaleSets
          name={name}
          resourceGroups={options.virtualmachinescalesets.resourceGroups}
          setDetails={showDetails}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'virtualmachinescalesets' &&
    options.virtualmachinescalesets &&
    options.virtualmachinescalesets.type === 'details' &&
    options.virtualmachinescalesets.resourceGroup &&
    options.virtualmachinescalesets.virtualMachineScaleSet
  ) {
    return (
      <PluginCard title={title} description={description}>
        <VMSSDetailsVirtualMachineScaleSets
          name={name}
          resourceGroup={options.virtualmachinescalesets.resourceGroup}
          virtualMachineScaleSet={options.virtualmachinescalesets.virtualMachineScaleSet}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'virtualmachinescalesets' &&
    options.virtualmachinescalesets &&
    options.virtualmachinescalesets.type === 'virtualMachines' &&
    options.virtualmachinescalesets.resourceGroup &&
    options.virtualmachinescalesets.virtualMachineScaleSet
  ) {
    return (
      <PluginCard title={title} description={description}>
        <VMSSDetailsVirtualMachines
          name={name}
          resourceGroup={options.virtualmachinescalesets.resourceGroup}
          virtualMachineScaleSet={options.virtualmachinescalesets.virtualMachineScaleSet}
        />
      </PluginCard>
    );
  }

  if (
    options?.type &&
    options?.type === 'virtualmachinescalesets' &&
    options.virtualmachinescalesets &&
    options.virtualmachinescalesets.type === 'metrics' &&
    options.virtualmachinescalesets.resourceGroup &&
    options.virtualmachinescalesets.virtualMachineScaleSet &&
    options.virtualmachinescalesets.metricNames &&
    options.virtualmachinescalesets.aggregationType &&
    times
  ) {
    return (
      <PluginCard title={title} description={description}>
        <Metric
          name={name}
          resourceGroup={options.virtualmachinescalesets.resourceGroup}
          provider={
            options.virtualmachinescalesets.virtualMachine
              ? providerVMSS +
                options.virtualmachinescalesets.virtualMachineScaleSet +
                '/virtualMachines/' +
                options.virtualmachinescalesets.virtualMachine.replace(
                  options.virtualmachinescalesets.virtualMachineScaleSet + '_',
                  '',
                )
              : providerVMSS + options.virtualmachinescalesets.virtualMachineScaleSet
          }
          metricNames={options.virtualmachinescalesets.metricNames}
          aggregationType={options.virtualmachinescalesets.aggregationType}
          times={times}
        />
      </PluginCard>
    );
  }

  return (
    <PluginOptionsMissing
      title={title}
      message="Options for Azure panel are missing or invalid"
      details="The panel doesn't contain the required options to render get the Azure data or the provided options are invalid."
      documentation="https://kobs.io/plugins/azure"
    />
  );
};

export default memo(Panel, (prevProps, nextProps) => {
  if (JSON.stringify(prevProps) === JSON.stringify(nextProps)) {
    return true;
  }

  return false;
});
