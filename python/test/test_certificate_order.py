"""
    SwissSign RA REST API

    See https://github.com/SwissSign-AG/RaApi/README.md  # noqa: E501

    The version of the OpenAPI document: 2.0.224
    Contact: ssc@swisssign.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import swisssign_ra_api.v2
from swisssign_ra_api.v2.model.certificate import Certificate
from swisssign_ra_api.v2.model.certificate_order_status import CertificateOrderStatus
globals()['Certificate'] = Certificate
globals()['CertificateOrderStatus'] = CertificateOrderStatus
from swisssign_ra_api.v2.model.certificate_order import CertificateOrder


class TestCertificateOrder(unittest.TestCase):
    """CertificateOrder unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testCertificateOrder(self):
        """Test CertificateOrder"""
        # FIXME: construct object with mandatory attributes with example values
        # model = CertificateOrder()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
