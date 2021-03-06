/* eslint-disable no-invalid-this */
import React from 'react';
import {shallow} from 'enzyme';
import {newTrain, noRequest, completeRequest} from 'test/TestData';
import Details from './Details';

describe('Phases', function() {
  beforeEach(function() {
    this.train = JSON.parse(JSON.stringify(newTrain));
  });

  it('Waits for train gracefully', function() {
    const wrapper = shallow(
      <Details
        train={null}
        request={noRequest}/>);
    expect(wrapper.debug()).toEqual(expect.stringContaining('Loading'));
  });

  it('Renders correctly', function() {
    const wrapper = shallow(
      <Details
        train={this.train}
        request={completeRequest}/>);
    expect(wrapper.debug()).toEqual(expect.stringContaining('Details'));
    expect(wrapper.debug()).not.toEqual(expect.stringContaining('Loading'));
  });
});
