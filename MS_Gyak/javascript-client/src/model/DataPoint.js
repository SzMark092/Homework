/*
 * Homework server
 * This is a simple server for my homework. It can make request towards a PSQL database, and get full tables of data from it.
 *
 * OpenAPI spec version: 2.0.0
 * Contact: szmgepesz@gmail.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.12
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.HomeworkServer) {
      root.HomeworkServer = {};
    }
    root.HomeworkServer.DataPoint = factory(root.HomeworkServer.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The DataPoint model module.
   * @module model/DataPoint
   * @version 2.0.0
   */

  /**
   * Constructs a new <code>DataPoint</code>.
   * @alias module:model/DataPoint
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>DataPoint</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/DataPoint} obj Optional instance to populate.
   * @return {module:model/DataPoint} The populated <code>DataPoint</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('ID'))
        obj.ID = ApiClient.convertToType(data['ID'], 'Number');
      if (data.hasOwnProperty('Timestamp'))
        obj.timestamp = ApiClient.convertToType(data['Timestamp'], 'Date');
      if (data.hasOwnProperty('DataPointDescription'))
        obj.dataPointDescription = ApiClient.convertToType(data['DataPointDescription'], Object);
      if (data.hasOwnProperty('Module'))
        obj.module = ApiClient.convertToType(data['Module'], Object);
      if (data.hasOwnProperty('Value'))
        obj.value = ApiClient.convertToType(data['Value'], 'Number');
    }
    return obj;
  }

  /**
   * @member {Number} ID
   */
  exports.prototype.ID = undefined;

  /**
   * @member {Date} timestamp
   */
  exports.prototype.timestamp = undefined;

  /**
   * @member {Object} dataPointDescription
   */
  exports.prototype.dataPointDescription = undefined;

  /**
   * @member {Object} module
   */
  exports.prototype.module = undefined;

  /**
   * @member {Number} value
   */
  exports.prototype.value = undefined;

  return exports;

}));