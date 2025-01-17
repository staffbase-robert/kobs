import { Button, ButtonVariant, SimpleListItem, Tooltip } from '@patternfly/react-core';
import { CopyIcon, LongArrowAltDownIcon, LongArrowAltUpIcon, TrashIcon } from '@patternfly/react-icons';
import React, { useState } from 'react';

export interface ILogsFieldsItemProps {
  index: number;
  length: number;
  field: string;
  selectField: (field: string) => void;
  changeFieldOrder: (oldIndex: number, newIndex: number) => void;
}

const LogsFieldsItem: React.FunctionComponent<ILogsFieldsItemProps> = ({
  index,
  length,
  field,
  selectField,
  changeFieldOrder,
}: ILogsFieldsItemProps) => {
  const [showActions, setShowActions] = useState<boolean>(false);

  const copyField = (field: string): void => {
    if (navigator.clipboard) {
      navigator.clipboard.writeText(field);
    }
  };

  return (
    <SimpleListItem key={index} isActive={false}>
      <div onMouseEnter={(): void => setShowActions(true)} onMouseLeave={(): void => setShowActions(false)}>
        {field}
        {showActions && (
          <div style={{ float: 'right' }}>
            {index !== length - 1 ? (
              <Tooltip content={<div>Move Down</div>}>
                <Button
                  style={{ padding: '0', paddingRight: '3px' }}
                  variant={ButtonVariant.plain}
                  aria-label="Move Down"
                  isSmall={true}
                  onClick={(): void => changeFieldOrder(index, index + 1)}
                >
                  <LongArrowAltDownIcon />
                </Button>
              </Tooltip>
            ) : null}

            {index !== 0 ? (
              <Tooltip content={<div>Move Up</div>}>
                <Button
                  style={{ padding: '0', paddingRight: '3px' }}
                  variant={ButtonVariant.plain}
                  aria-label="Move Up"
                  isSmall={true}
                  onClick={(): void => changeFieldOrder(index, index - 1)}
                >
                  <LongArrowAltUpIcon />
                </Button>
              </Tooltip>
            ) : null}

            <Tooltip content={<div>Copy</div>}>
              <Button
                style={{ padding: '0', paddingRight: '3px' }}
                variant={ButtonVariant.plain}
                aria-label="Copy"
                isSmall={true}
                onClick={(): void => copyField(field)}
              >
                <CopyIcon />
              </Button>
            </Tooltip>

            <Tooltip content={<div>Remove</div>}>
              <Button
                style={{ padding: '0', paddingRight: '3px' }}
                variant={ButtonVariant.plain}
                aria-label="Remove"
                isSmall={true}
                onClick={(): void => selectField(field)}
              >
                <TrashIcon />
              </Button>
            </Tooltip>
          </div>
        )}
      </div>
    </SimpleListItem>
  );
};

export default LogsFieldsItem;
