import React from 'react';

import {shallow, ShallowWrapper} from 'enzyme';

import {ModalSubtitleAndError} from '@brightscout/mattermost-ui-library';

import {SubscriptionType} from 'plugin_constants';

import SubscriptionTypePanel from './subscriptionTypePanel';

const mockOnContinue = jest.fn();
const mockOnBack = jest.fn();
const mockSetSubscriptionType = jest.fn();

const subscriptionTypePanelProps = {
    className: 'mockClassName',
    error: 'mockError',
    onBack: mockOnBack,
    onContinue: mockOnContinue,
    actionBtnDisabled: true,
    requiredFieldValidationErr: true,
    subscriptionType: SubscriptionType.RECORD,
    setSubscriptionType: mockSetSubscriptionType,
};

describe('Subscription Type Panel', () => {
    let component: ShallowWrapper<any, Readonly<{}>, React.Component<{}, {}, any>>;

    beforeEach(() => {
        component = shallow(<SubscriptionTypePanel {...subscriptionTypePanelProps}/>);
    });

    it('Should render correctly', () => {
        expect(component).toMatchSnapshot();
    });

    it('Should apply the passed className prop', () => {
        expect(component.hasClass(subscriptionTypePanelProps.className)).toBeTruthy();
    });

    it('Should render the subscription type panel body correctly', () => {
        expect(component.find('Dropdown')).toHaveLength(1);
        expect(component.find('ModalSubtitleAndError')).toHaveLength(1);
        expect(component.find('ModalFooter')).toHaveLength(1);
    });

    it('Should render the error correctly', () => {
        expect(component.contains(
            <ModalSubtitleAndError error={subscriptionTypePanelProps.error}/>,
        )).toBeTruthy();
    });

    it('Should fire change event when dropdown value is changed', () => {
        const changeDropdown = (changeNumber: number) => {
            component.find('Dropdown').simulate('change');
            expect(subscriptionTypePanelProps.setSubscriptionType).toHaveBeenCalledTimes(changeNumber);
        };

        // Click the checkbox
        changeDropdown(1);
        changeDropdown(2);
    });
});
