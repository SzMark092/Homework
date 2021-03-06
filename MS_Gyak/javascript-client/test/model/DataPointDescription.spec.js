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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.HomeworkServer);
  }
}(this, function(expect, HomeworkServer) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('DataPointDescription', function() {
      beforeEach(function() {
        instance = new HomeworkServer.DataPointDescription();
      });

      it('should create an instance of DataPointDescription', function() {
        // TODO: update the code to test DataPointDescription
        expect(instance).to.be.a(HomeworkServer.DataPointDescription);
      });

      it('should have the property ID (base name: "ID")', function() {
        // TODO: update the code to test the property ID
        expect(instance).to.have.property('ID');
        // expect(instance.ID).to.be(expectedValueLiteral);
      });

      it('should have the property name (base name: "Name")', function() {
        // TODO: update the code to test the property name
        expect(instance).to.have.property('name');
        // expect(instance.name).to.be(expectedValueLiteral);
      });

      it('should have the property title (base name: "Title")', function() {
        // TODO: update the code to test the property title
        expect(instance).to.have.property('title');
        // expect(instance.title).to.be(expectedValueLiteral);
      });

      it('should have the property max (base name: "Max")', function() {
        // TODO: update the code to test the property max
        expect(instance).to.have.property('max');
        // expect(instance.max).to.be(expectedValueLiteral);
      });

      it('should have the property min (base name: "Min")', function() {
        // TODO: update the code to test the property min
        expect(instance).to.have.property('min');
        // expect(instance.min).to.be(expectedValueLiteral);
      });

    });
  });

}));
