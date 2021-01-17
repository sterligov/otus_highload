import React from 'react';

export default class Error extends React.Component {
    render() {
        const error = this.props.error;
        return (
            <div>
                <div className="d-flex justify-content-center">
                    <h1>
                        <strong>
                            {error !== undefined &&
                            <p>Произошла ошибка {error.status}</p>
                            }
                            {error === undefined &&
                            <p>Ошибка! Кажется что-то пошло не так. Попробуйте позже.</p>
                            }
                        </strong>
                    </h1>
                </div>
            </div>
        );
    }
}