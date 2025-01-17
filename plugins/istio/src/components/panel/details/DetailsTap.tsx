import {
  Card,
  CardBody,
  DrawerActions,
  DrawerCloseButton,
  DrawerHead,
  DrawerPanelBody,
  DrawerPanelContent,
} from '@patternfly/react-core';
import { TableComposable, TableVariant, Tbody, Td, Tr } from '@patternfly/react-table';
import React from 'react';

import { formatTime, getDirection } from '../../../utils/helpers';
import { ILogLine } from '../../../utils/interfaces';
import { Title } from '@kobsio/plugin-core';

interface IDetailsTapProps {
  line: ILogLine;
  close: () => void;
}

const DetailsTap: React.FunctionComponent<IDetailsTapProps> = ({ line, close }: IDetailsTapProps) => {
  return (
    <DrawerPanelContent minSize="50%">
      <DrawerHead>
        <Title
          title={`${
            line.hasOwnProperty('content.upstream_cluster') ? getDirection(line['content.upstream_cluster']) : '-'
          }: ${line.hasOwnProperty('content.authority') ? line['content.authority'] : '-'}`}
          subtitle={formatTime(line['time'])}
          size="lg"
        />
        <DrawerActions style={{ padding: 0 }}>
          <DrawerCloseButton onClose={close} />
        </DrawerActions>
      </DrawerHead>

      <DrawerPanelBody>
        <Card isCompact={true}>
          <CardBody>
            <TableComposable aria-label="Details" variant={TableVariant.compact} borders={false}>
              <Tbody>
                {Object.keys(line).map((key) => (
                  <Tr key={key}>
                    <Td noPadding={true} dataLabel="Key">
                      <b>{key}</b>
                    </Td>
                    <Td className="pf-u-text-wrap pf-u-text-break-word" noPadding={true} dataLabel="Value">
                      <div style={{ whiteSpace: 'pre-wrap' }}>{line[key]}</div>
                    </Td>
                  </Tr>
                ))}
              </Tbody>
            </TableComposable>
          </CardBody>
        </Card>
      </DrawerPanelBody>
    </DrawerPanelContent>
  );
};

export default DetailsTap;
