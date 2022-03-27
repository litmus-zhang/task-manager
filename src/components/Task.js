import React from 'react'
import PropTypes from 'prop-types'

export default function Task({task: {id, title, state}, onArchiveTask, onPinTask
}) {
  return (
      <div className={`list-item ${state}`}>
          <label className='checkbox'>
              <input
                  type='checkbox'
                  name='checked'
                  defaultChecked={state === "TASK_ARCHIVED"}
                  disabled={true}
              />
              <span
                  className='checkbox-custom'
                  onClick={() => onArchiveTask(id)}
                  id={`archiveTask-${id}`}
              />
          </label>
          <div className='title'>
              <input type='text' value={title} readOnly={true} placeholder="Input Title" />
          </div>

          <div className='actions' onClick={event => event.stopPropagation()}>
              {
                  state !== 'TASK_ARCHIVED' && (
                      // eslint-disable-next-line jsx-a11y/anchor-is-valid
                      <a onClick={()=> onPinTask(id)}>
                          <span className={ `icon-star` } id={`pinTask-${id}`} />
                      </a>
                  )
              }

          </div>
    </div>
  )
}

Task.prototype = {
    /**Composition of the task */
    task: PropTypes.shape({
        //ID of the task
        id: PropTypes.string.isRequired,
         //Title of the task
        title: PropTypes.string.isRequired,

        //State of the task
        state: PropTypes.string.isRequired,
    }),
    //Events to change the task to archived
    onArchiveTask: PropTypes.func,
    //Events to change the task to pinned
    onPinTask: PropTypes.func,

}