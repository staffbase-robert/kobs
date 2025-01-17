import React, { useState } from 'react';
import { Tab, TabTitleText, Tabs } from '@patternfly/react-core';
import { TableComposable, TableVariant, Tbody } from '@patternfly/react-table';

import { Editor } from '@kobsio/plugin-core';
import { IDocument } from '../../utils/interfaces';
import LogsDocumentDetailsRow from './LogsDocumentDetailsRow';

export interface ILogsDocumentDetailsProps {
  document: IDocument;
  addFilter?: (filter: string) => void;
  selectField?: (field: string) => void;
}

const LogsDocumentDetails: React.FunctionComponent<ILogsDocumentDetailsProps> = ({
  document,
  addFilter,
  selectField,
}: ILogsDocumentDetailsProps) => {
  const [activeTab, setActiveTab] = useState<string>('table');

  return (
    <Tabs
      activeKey={activeTab}
      onSelect={(event, tabIndex): void => setActiveTab(tabIndex.toString())}
      isFilled={true}
      mountOnEnter={true}
    >
      <Tab eventKey="table" title={<TabTitleText>Table</TabTitleText>}>
        <div style={{ maxWidth: '100%', overflowX: 'scroll', padding: '24px 0px' }}>
          <TableComposable aria-label="Details" variant={TableVariant.compact} borders={false}>
            <Tbody>
              {Object.keys(document).map((key) => (
                <LogsDocumentDetailsRow
                  key={key}
                  documentKey={key}
                  documentValue={document[key]}
                  addFilter={addFilter}
                  selectField={selectField}
                />
              ))}
            </Tbody>
          </TableComposable>
        </div>
      </Tab>
      <Tab eventKey="json" title={<TabTitleText>JSON</TabTitleText>}>
        <div style={{ maxWidth: '100%', overflowX: 'scroll', padding: '24px 0px' }}>
          <Editor value={JSON.stringify(document, null, 2)} mode="json" readOnly={true} />
        </div>
      </Tab>
    </Tabs>
  );
};

export default LogsDocumentDetails;
