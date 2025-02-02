import {
  Alert,
  AlertActionCloseButton,
  AlertGroup,
  AlertVariant,
  CardActions,
  Dropdown,
  DropdownItem,
  KebabToggle,
} from '@patternfly/react-core';
import React, { useState } from 'react';

interface IDetailsContainerGroupActionsProps {
  name: string;
  resourceGroup: string;
  containerGroup: string;
  isPanelAction: boolean;
}

const DetailsContainerGroupActions: React.FunctionComponent<IDetailsContainerGroupActionsProps> = ({
  name,
  resourceGroup,
  containerGroup,
  isPanelAction,
}: IDetailsContainerGroupActionsProps) => {
  const [showDropdown, setShowDropdown] = useState<boolean>(false);
  const [alert, setAlert] = useState<{ title: string; variant: AlertVariant } | undefined>(undefined);

  const restart = async (): Promise<void> => {
    try {
      const response = await fetch(
        `/api/plugins/azure/${name}/containerinstances/containergroup/restart?resourceGroup=${resourceGroup}&containerGroup=${containerGroup}`,
        {
          method: 'put',
        },
      );
      const json = await response.json();

      if (response.status >= 200 && response.status < 300) {
        setShowDropdown(false);
        setAlert({ title: `${containerGroup} (${resourceGroup}) was restarted`, variant: AlertVariant.success });
      } else {
        if (json.error) {
          throw new Error(json.error);
        } else {
          throw new Error('An unknown error occured');
        }
      }
    } catch (err) {
      setShowDropdown(false);
      setAlert({ title: err.message, variant: AlertVariant.danger });
    }
  };

  const dropdown = (
    <Dropdown
      className={isPanelAction ? 'pf-c-drawer__close' : undefined}
      toggle={<KebabToggle onToggle={(): void => setShowDropdown(!showDropdown)} />}
      isOpen={showDropdown}
      isPlain={true}
      position="right"
      dropdownItems={[
        <DropdownItem key={0} onClick={restart}>
          Restart
        </DropdownItem>,
      ]}
    />
  );

  if (isPanelAction) {
    return <CardActions>{dropdown}</CardActions>;
  }

  return (
    <React.Fragment>
      {alert ? (
        <AlertGroup isToast={true}>
          <Alert
            isLiveRegion={true}
            variant={alert.variant}
            title={alert.title}
            actionClose={<AlertActionCloseButton onClick={(): void => setAlert(undefined)} />}
          />
        </AlertGroup>
      ) : null}

      {dropdown}
    </React.Fragment>
  );
};

export default DetailsContainerGroupActions;
