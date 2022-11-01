import {RecordType} from 'plugin_constants';

import reducer, {showModal, hideModal, UpdateStateModalState} from 'reducers/updateStateModal';

const previousState: UpdateStateModalState = {
    open: false,
}

test('should change the state of open to "true" and data value to payload', () => {
    const payload: CommentAndStateModalData = {
        recordType: RecordType.INCIDENT,
        recordId: "mockId",
    };
    expect(reducer(previousState, showModal(payload))).toEqual(
        {open: true, data: payload},
    );
});

test('should change the state of open to "false"', () => {
    previousState.open = true;
    expect(reducer(previousState, hideModal())).toEqual(
        {open: false},
    );
});
