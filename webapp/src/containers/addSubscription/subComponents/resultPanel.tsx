import React, {forwardRef} from 'react';

type ResultPanelProps = {
    iconClass?: string | null;
    header?: string | null;
    className?: string;
    primaryBtnText?: string;
    secondaryBtnText?: string;
    onPrimaryBtnClick?: () => void | null;
    onSecondaryBtnClick?: () => void | null;
};

const ResultPanel = forwardRef<HTMLDivElement, ResultPanelProps>(({
    header,
    className = '',
    primaryBtnText,
    secondaryBtnText,
    onPrimaryBtnClick,
    onSecondaryBtnClick,
    iconClass,
}: ResultPanelProps, resultPanelRef): JSX.Element => (
    <div
        className={`modal__body modal-body d-flex align-items-center justify-content-center flex-column secondary-panel ${className}`}
        ref={resultPanelRef}
    >
        <>
            <i className={`fa result-panel-icon ${iconClass ?? 'fa-check-circle-o'}`}/>
            <h2 className='result-panel-text'>{header ?? 'Subscription added successfully!'}</h2>
            {onPrimaryBtnClick && (
                <button
                    className='btn btn-primary'
                    onClick={onPrimaryBtnClick}
                >
                    {primaryBtnText ?? 'Add Another Subscription'}
                </button>
            )}
            {onSecondaryBtnClick && (
                <button
                    className='btn btn-link result-panel-close-btn'
                    onClick={onSecondaryBtnClick}
                >
                    {secondaryBtnText ?? 'Close'}
                </button>
            )}
        </>
    </div>
));

export default ResultPanel;
